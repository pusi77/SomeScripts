#!/usr/bin/python3

import paramiko
import sys

username = 'CHANGEME'
ip_address = 'CHANGEME'


usage = 'Usage: ./phdis.py password interval\ne.g.: To stop Pi-Hole for 10 min -> "./phdis.py mypassword 10"'


def check_args():
    if len(sys.argv) != 3:
        if len(sys.argv) == 2 and sys.argv[1] == '-h':
            return 1
        else:
            raise SyntaxError(f'Invalid arguments.\n{usage}')
    int(sys.argv[2])  # raise a ValueError if it's not an int
    return 0


def check_changeme():
    if (username == 'CHANGEME') and (ip_address == 'CHANGEME'):
        return 1
    return 0


def connect_and_run(interval, password):
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    ssh.load_system_host_keys()

    try:
        ssh.connect(ip_address, username=username, password=password)
        print('Connected')
        command = f'pihole disable {interval}m'
        ssh_stdin, ssh_stdout, ssh_stderr = ssh.exec_command(command)
        print('Response: ')
        for item in ssh_stdout:
            print(item, end='')
        ssh.close()
        print('Connection closed')
    except Exception as e:
        print('Eccezione: {}', e)



if __name__ == '__main__':
    if(check_changeme() != 0):
        print('Please change username and ip_address values hardcoded in the script.')
        sys.exit()

    try:
        f'Second argument must be an integer.\n{usage}'
        if check_args():
            sys.exit()
    except SyntaxError as se:
        print(se)
        sys.exit()
    except ValueError as ve:
        print(f'Second argument must be an integer.\n{usage}')
        sys.exit()

    try:
        connect_and_run(sys.argv[2], sys.argv[1])
    except Exception as e:
        #connect raises BadHostKeyException, AuthenticationException
        #               SSHException and socket.error
        print(e)