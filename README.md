# go-logr-issue-rg0

Sample issue for training puposes.

Run the main file via `go run main.go` and you should see the following log lines:

```
{"level":"dpanic","ts":1739977744.980593,"caller":"go-logr-issue-rg0/main.go:28","msg":"odd number of arguments passed as key-value pairs for logging","event":"GetOrganization","project-id":"project-id-1","ignored key":null,"stacktrace":"main.withProjectLogger\n\t/Users/gogolokr/workspace/github.com/gogolok/go-logr-issue-rg0/main.go:28\nmain.withProjectOrgLogger\n\t/Users/gogolokr/workspace/github.com/gogolok/go-logr-issue-rg0/main.go:23\nmain.main\n\t/Users/gogolokr/workspace/github.com/gogolok/go-logr-issue-rg0/main.go:15\nruntime.main\n\t/opt/homebrew/Cellar/go/1.24.0/libexec/src/runtime/proc.go:283"}
{"level":"info","ts":1739977744.9809742,"caller":"go-logr-issue-rg0/main.go:16","msg":"event received","event":"GetOrganization","project-id":"project-id-1","organization-id":"org-id-3"}
DONE
```

# Solution

```diff
diff --git a/main.go b/main.go
index e06369e..4646f98 100644
--- a/main.go
+++ b/main.go
@@ -20,7 +20,7 @@ func main() {

 func withProjectOrgLogger(logger logr.Logger, event string, projectId string, organizationId string, keysAndValues ...any) logr.Logger {
        const logKeyOrganization = "organization-id"
-       return withProjectLogger(logger, event, projectId, logKeyOrganization, organizationId, keysAndValues)
+       return withProjectLogger(logger, event, projectId, logKeyOrganization, organizationId).WithValues(keysAndValues...)
 }

 func withProjectLogger(logger logr.Logger, event string, projectId string, keysAndValues ...any) logr.Logger {
```
