#!/usr/bin/env bash
set -euo pipefail

dotnet tool install --global MarkdownSnippets.Tool
mdsnippets .