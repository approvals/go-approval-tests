#!/usr/bin/env python3
import csv
import os
import re
import sys


def screaming_snake_to_pascal(name: str) -> str:
    return "".join(word.capitalize() for word in name.split("_"))


def parse_arguments_template(arguments: str) -> list[str]:
    if not arguments or not arguments.strip():
        return []
    return arguments.split()


def build_go_args_from_template(arguments: str) -> str:
    parts = parse_arguments_template(arguments)
    go_parts = []
    next_var = ["received"]

    def substitute_token(token: str) -> str:
        if "%s" not in token:
            return f'"{go_escape_string(token)}"'
        segments = token.split("%s")
        vars_used = []
        for _ in range(len(segments) - 1):
            vars_used.append(next_var[0])
            next_var[0] = "approved" if next_var[0] == "received" else "received"
        bits = [f'"{go_escape_string(segments[0])}"']
        for i, v in enumerate(vars_used):
            bits.append(v)
            if segments[i + 1]:
                bits.append(f'"{go_escape_string(segments[i + 1])}"')
        return " + ".join(bits)

    for p in parts:
        if p == "%s":
            var = next_var[0]
            next_var[0] = "approved" if next_var[0] == "received" else "received"
            go_parts.append(var)
        else:
            go_parts.append(substitute_token(p))
    if not go_parts:
        return "[]string{received, approved}"
    return "[]string{" + ", ".join(go_parts) + "}"


def go_escape_string(s: str) -> str:
    return s.replace("\\", "\\\\").replace('"', '\\"')

def path_go_literal(path: str, os_name: str) -> str:
    path = path.strip()
    if os_name == "Windows" and "{ProgramFiles}" in path:
        return f'expandProgramFiles("{go_escape_string(path)}")'
    path = path.replace("\\", "/")
    return f'"{path}"'


def os_to_goos(os_name: str) -> str:
    m = {"Mac": "goosDarwin", "Windows": "goosWindows", "Linux": "goosLinux"}
    return m[os_name]


def generate_reporter_struct(row: dict) -> str:
    name = row["name"]
    path = row["path"].strip()
    arguments = row["arguments"].strip()
    os_name = row["os"]
    struct_name = screaming_snake_to_pascal(name) + os_name
    struct_name = struct_name[0].lower() + struct_name[1:]
    goos = os_to_goos(os_name)
    path_expr = path_go_literal(path, os_name)
    args_expr = build_go_args_from_template(arguments)
    return f'''type {struct_name} struct{{}}

func New{screaming_snake_to_pascal(name)}{os_name}Reporter() Reporter {{
	return &{struct_name}{{}}
}}

func (s *{struct_name}) Report(approved, received string) bool {{
	if runtime.GOOS != {goos} {{
		return false
	}}
	programName := {path_expr}
	args := {args_expr}
	return launchProgram(programName, approved, args...)
}}
'''


def generate_per_os_aggregator(os_name: str, reporter_rows: list[dict]) -> str:
    goos = os_to_goos(os_name)
    struct_name = "diffToolOn" + os_name
    constructor_name = "NewDiffToolOn" + os_name + "Reporter"
    new_calls = []
    for row in reporter_rows:
        name = row["name"]
        new_calls.append(f"New{screaming_snake_to_pascal(name)}{os_name}Reporter()")
    reporters_list = ",\n\t\t".join(new_calls)
    return f'''type {struct_name} struct{{}}

func {constructor_name}() Reporter {{
	return NewFirstWorkingReporter(
		{reporters_list},
	)
}}

func (s *{struct_name}) Report(approved, received string) bool {{
	if runtime.GOOS != {goos} {{
		return false
	}}
	return NewFirstWorkingReporter(
		{reporters_list},
	).Report(approved, received)
}}
'''


def generate_group_aggregator(group_name: str, reporter_rows: list[dict]) -> str:
    group_pascal = screaming_snake_to_pascal(group_name)
    struct_name = group_pascal[0].lower() + group_pascal[1:] + "Group"
    constructor_name = "New" + group_pascal + "GroupReporter"
    new_calls = []
    for row in reporter_rows:
        name = row["name"]
        os_name = row["os"]
        new_calls.append(f"New{screaming_snake_to_pascal(name)}{os_name}Reporter()")
    reporters_list = ",\n\t\t".join(new_calls)
    return f'''type {struct_name} struct{{}}

func {constructor_name}() Reporter {{
	return NewFirstWorkingReporter(
		{reporters_list},
	)
}}

func (s *{struct_name}) Report(approved, received string) bool {{
	return NewFirstWorkingReporter(
		{reporters_list},
	).Report(approved, received)
}}
'''


def main() -> None:
    repo_root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    csv_path = os.path.join(repo_root, "DiffTools", "diff_reporters.csv")
    out_path = os.path.join(repo_root, "reporters", "diff_reporters_generated.go")

    rows = []
    with open(csv_path, newline="", encoding="utf-8") as f:
        reader = csv.DictReader(f)
        for row in reader:
            if not any(row.values()):
                continue
            rows.append(row)

    by_os = {"Mac": [], "Windows": [], "Linux": []}
    by_group = {}

    sections = []
    sections.append('''package reporters

import (
	"os"
	"runtime"
	"strings"
)

func expandProgramFiles(path string) string {
	return strings.ReplaceAll(strings.Replace(path, "{ProgramFiles}", os.Getenv("ProgramFiles"), 1), "\\\\", "/")
}

''')

    for row in rows:
        sections.append(generate_reporter_struct(row))
        by_os[row["os"]].append(row)
        group = row.get("group_name", "").strip()
        if group:
            by_group.setdefault(group, []).append(row)

    for os_name in ("Mac", "Windows", "Linux"):
        if by_os[os_name]:
            sections.append(generate_per_os_aggregator(os_name, by_os[os_name]))

    for group_name in sorted(by_group.keys()):
        sections.append(generate_group_aggregator(group_name, by_group[group_name]))

    content = "\n".join(sections)
    with open(out_path, "w", encoding="utf-8") as f:
        f.write(content)

    print(out_path, file=sys.stderr)


if __name__ == "__main__":
    main()
