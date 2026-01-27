---
name: gitlab
description: >
  Required skill to perform GitLab MR operations, manage discussion
  threads, and check CI pipeline status using the glab CLI.
user-invocable: false
---

# GitLab MR Interactions

Concrete `glab` commands for common GitLab merge request operations.

Note: `:fullpath` is a `glab api` placeholder that resolves to the current
project's full path (e.g., `group/project`). Other placeholders like
`<mr-iid>` are literal values you supply.

<action name="create-mr">

```bash
glab mr create --title "<title>" --target-branch <target> --description "<body>"
```

Omit `--target-branch` to target the repository's default branch.

</action>

<action name="update-mr">

Change the target branch:

```bash
glab mr update <mr-iid> --target-branch <new-target>
```

Update title or description:

```bash
glab mr update <mr-iid> --title "<new-title>"
glab mr update <mr-iid> --description "<new-description>"
```

</action>

<action name="view-mr">

```bash
glab mr view <mr-iid>
glab mr view <mr-iid> --comments
```

</action>

<action name="discussions">

`glab mr note` creates top-level comments only. Use `glab api` to reply in
existing threads.

<action name="top-level-comment">

```bash
glab mr note <mr-iid> -m "<message>"
```

</action>

<action name="list-discussions">

```bash
glab api projects/:fullpath/merge_requests/<mr-iid>/discussions
```

This returns a JSON array. Each discussion has an `id` field (the
`discussion_id`) and a `notes` array containing the individual comments.

</action>

<action name="reply-in-thread">

```bash
glab api --method POST \
  "projects/:fullpath/merge_requests/<mr-iid>/discussions/<discussion_id>/notes" \
  -f body="<reply text>"
```

Replace `<discussion_id>` with the discussion's `id` from the listing above.

<constraint>
Thread resolution triggers GitLab workflows and notifications. Only users
should resolve threads — NEVER resolve them programmatically.
</constraint>

</action>

</action>

<action name="ci-pipelines">

<action name="pipeline-status">

```bash
glab ci status
glab ci status --branch <branch>
```

</action>

<action name="job-logs">

By job name:

```bash
glab ci trace <job-name>
glab ci trace <job-name> --branch <branch>
```

By pipeline ID:

```bash
glab ci trace --pipeline-id <id> <job-name>
```

</action>

<action name="list-jobs">

```bash
glab ci list
glab ci list --branch <branch>
```

</action>

</action>
