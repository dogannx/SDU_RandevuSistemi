pipeline {
    agent any

    options {
        timestamps()
        timeout(time: 20, unit: 'MINUTES')
    }

    stages {
        stage('Ortam Bilgisi') {
            steps {
                sh 'go version'
                sh 'node --version'
                sh 'npm --version'
            }
        }

        stage('Build & Test') {
            parallel {
                stage('Backend (Go)') {
                    steps {
                        dir('backend') {
                            sh 'go mod download'
                            sh 'go vet ./...'
                            sh 'go build ./...'
                            sh 'go test ./...'
                        }
                    }
                }

                stage('Frontend (Vue)') {
                    steps {
                        dir('frontend') {
                            sh 'npm ci'
                            sh 'npm run build'
                        }
                    }
                }
            }
        }
    }

    post {
        success {
            echo 'CI başarılı — backend testleri ve frontend build geçti.'
        }
        failure {
            echo 'CI başarısız — logları kontrol et.'
        }
    }
}
