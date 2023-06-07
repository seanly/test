import jenkins.model.Jenkins
import com.cloudbees.hudson.plugins.folder.AbstractFolder

// 获取 Jenkins 实例
Jenkins jenkins = Jenkins.getInstance()

// 获取目录 "demo"
def demoFolder = jenkins.getItemByFullName("demo")

// 如果目录存在，则获取该目录下的所有 job
if (demoFolder instanceof AbstractFolder) {
    List<String> jobNames = demoFolder.getAllItems().collect { item ->
        if (!(item instanceof AbstractFolder)) {
            return item.fullName
        }
    }.findAll { it != null }

    // 删除非运行中的 job
    for (String jobName : jobNames) {
        // 获取 job 实例
        def job = jenkins.getItemByFullName(jobName)

        // 检查 job 是否存在并且不在运行中
        if (job && !job.isBuilding()) {
            jenkins.deleteJob(jobName)
            println("Deleted job: ${jobName}")
        } else if (job) {
            println("Skipping job: ${jobName} (Job is currently running)")
        } else {
            println("Skipping job: ${jobName} (Job not found)")
        }
    }
} else {
    println("Folder 'demo' not found")
}
