# Process for Fixing Bugs

When fixing a bug reported in an issue or pull request:

If you don't have the URL of the issue or pull request, ask for it.

1.  **Understand the Bug**:
    *   Thoroughly read the issue/PR description and any related comments to understand the problem, expected behavior, and actual behavior.
    *  Read the pull request via the github cli: gh pr view <number> --json body

2.  **Write a Failing Test**:
    *   Create a new test case or modify an existing one in the relevant test suite
    *   The test should specifically target the scenario described in the bug report and fail with the current codebase, clearly demonstrating the bug (e.g., throwing the specific exception, producing incorrect output).

3.  **Implement the Fix**:
    *   Modify the necessary source code file(s) to correct the bug.
    *   Focus on addressing the root cause identified by the failing test.

4.  **Verify the Fix**:
    *   Run the tests
    *   Confirm that the previously failing test now passes.
    *   Ensure **all** other tests continue to pass, verifying no regressions were introduced.

5.  **Commit the Changes**:
    *   Use the Arlo Commit Notation (`<Risk Symbol> B <Title>`). Typically, the risk symbol will be `-` (Tested) if you wrote a unit test.
    *   Reference the fixed issue number in the commit message body (e.g., `Fixes #<issue_number>`).
    *   Include any necessary co-authors (refer to `ArloCommitNotation.process.md`).
    *   Use the commit script: `scripts/commit.sh "<commit_message>"`