import jenkins.model.Jenkins

// 获取 Jenkins 实例
Jenkins jenkins = Jenkins.getInstance()

// 获取所有的 jobs
List<String> jobNames = jenkins.getAllItems().collect { item ->
    // 仅处理非目录类型的 job
    if (!(item instanceof com.cloudbees.hudson.plugins.folder.AbstractFolder)) {
        return item.fullName
    } else {
        return null
    }
}.findAll { it != null }

// 删除非运行中的 job
for (String jobName : jobNames) {
    // 获取 job 实例
    def job = jenkins.getItemByFullName(jobName)

    // 检查 job 是否存在并且不在运行中
    if (job && !job.isBuilding()) {
        //jenkins.deleteJob(jobName)
        println("Deleted job: ${jobName}")
    } else if (job) {
        println("Skipping job: ${jobName} (Job is currently running)")
    } else {
        println("Skipping job: ${jobName} (Job not found)")
    }
}
