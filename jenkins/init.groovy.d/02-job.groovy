#!groovy
import jenkins.model.*
import org.jenkinsci.plugins.workflow.job.WorkflowJob
import org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition

def jobName = 'SDU_RandevuSistemi'
def instance = Jenkins.get()

def pipelineScript = '''\
node {
    stage('Checkout') {
        git url: 'https://github.com/dogannx/SDU_RandevuSistemi.git', branch: 'feature/backend-redis'
    }
    stage('Ortam Bilgisi') {
        sh 'go version'
        sh 'node --version'
        sh 'npm --version'
    }
    stage('Build & Test') {
        parallel(
            'Backend (Go)': {
                dir('backend') {
                    sh 'go mod download'
                    sh 'go vet ./...'
                    sh 'go build ./...'
                    sh 'go test ./...'
                }
            },
            'Frontend (Vue)': {
                dir('frontend') {
                    sh 'npm ci'
                    sh 'npm run build'
                }
            }
        )
    }
}
'''

if (instance.getItem(jobName) == null) {
    def job = instance.createProject(WorkflowJob.class, jobName)
    job.setDefinition(new CpsFlowDefinition(pipelineScript, true))
    job.save()
    println ">>> Pipeline job '${jobName}' oluşturuldu"
} else {
    println ">>> Pipeline job '${jobName}' zaten mevcut"
}
