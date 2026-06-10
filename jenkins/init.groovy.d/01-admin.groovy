#!groovy
import jenkins.model.*
import hudson.security.*
import jenkins.install.InstallState

def instance = Jenkins.get()

if (!(instance.getSecurityRealm() instanceof HudsonPrivateSecurityRealm)) {
    def realm = new HudsonPrivateSecurityRealm(false)
    realm.createAccount('admin', 'admin')
    instance.setSecurityRealm(realm)

    def strategy = new FullControlOnceLoggedInAuthorizationStrategy()
    strategy.setAllowAnonymousRead(false)
    instance.setAuthorizationStrategy(strategy)

    instance.setInstallState(InstallState.INITIAL_SETUP_COMPLETED)
    instance.save()

    println '>>> Admin user "admin" oluşturuldu (parola: admin)'
}
