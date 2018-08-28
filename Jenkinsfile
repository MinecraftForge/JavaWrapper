node {
    def root = tool name:'Go1.11' type: 'go'
    ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/src/github.com/MinecraftForge/JavaWrapper") {
        withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/", "PATH+GO=${root}/bin"]) {
            env.PATH="${GOPATH}/bin:$PATH"

            stages {
                stage('fetch') {
                    git(url: 'https://github.com/MinecraftForge/JavaWrapper.git')
                }

                stage('prebuild') {
                    sh 'go version'
                }

                stage('build') {
                    sh 'make compileWin'
                }
            }
            
        }
    }
}