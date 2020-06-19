### Job Configuration:

pkg/job

The job package wraps the pkg/filetransfer, pkg/trigger and pkg/execute to execute tasks.

This application transfers files by executing jobs, the jobs defined pre and post transfer logic for determining how and whether the transfer should take place.
Each job consists of Actions in order to do so. These are the following actions 
1. Trigger: This handles the start of the job.
  - [ ] Manual Request
  - [ ] Cron
  - [ ] Inotify
2. PreExecutor:
  This is for executing something before the start of a job. eg: send http request or whitelist connection or run a check.
  - [ ] Transfer Check - checks that it's okay to transfer from the directory defined in the job.
3. Matcher: This matches files based on timestamp and/or pattern. Need to change this from switch-case to if-else.
  - [ ] Transfer Check - checks that it's okay to transfer from the directory defined in the job.
  - [ ] Matcher
4. FileCheck: This runs a check on individual files that are about to be transferred.
  - [ ] OL - OL or Overwrite Logic handles the logic for files that are about to be overwritten.
5. Transfer: 
  - [ ] Pull
  - [ ] Push
  - [ ] PullExternal
  - [ ] PushExternal
6. PostExecutor: This is for Executing something at the end of a job.
  - [ ] CmdExecutor - for executing a command.

* Trigger
* Source
- NodeID -
- SPath
- SFile
- SCondition
- SExecutor
* Destination - 
