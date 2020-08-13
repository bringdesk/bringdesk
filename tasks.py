
from invoke import task

@task
def depend(c):
    c.run('apt-get install libcairo2-dev -y', pty=True, echo=True)
    c.run('npm install -g node-gyp', pty=True, echo=True)
    c.run('npm install', pty=True, echo=True)

@task
def clean(c):
    pass

@task(default=True, pre=[ clean ])
def build(c):
    c.run('cd bringdesk_screen && node-gyp configure', pty=True, echo=True)
    c.run('cd bringdesk_screen && node-gyp build', pty=True, echo=True)

@task
def check(c):
    c.run('ninja -C builddir test', pty=True)

@task
def install(c):
    c.run('install -m 0644 ./contrib/bringdesk.service /etc/systemd/system/bringdesk.service', pty=True)
    c.run('systemctl daemon-reload', pty=True)

@task
def stop(c):
    c.run('systemctl stop bringdesk.service', pty=True)

@task
def status(c):
    c.run('systemctl status bringdesk.service', pty=True)

@task
def start(c):
    c.run('systemctl start bringdesk.service', pty=True)

@task
def restart(c):
    c.run('systemctl restart bringdesk.service', pty=True)
