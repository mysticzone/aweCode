"""
-------------------------------------------------
   FileName：     audit_agent
   Date：         2022/12/6
-------------------------------------------------
"""
import fcntl
import socket
import logging
import sys
import struct
import http.client
import urllib.parse

NET_DRIVER = ""
CONNECTION_TIMEOUT = 100
OMSERVER_ADDRESS = ""

socket.setdefaulttimeout(CONNECTION_TIMEOUT)
logging.basicConfig(level=logging.DEBUG,
                    format='%(asctime)s [%(levelname)s] %(message)s',
                    filename=sys.path[0] + "./omsys.log",
                    filemode="a")

if len(sys.argv) < 6:
    logging.error("Not configured in /etc/profile !")
    sys.exit()


def get_local_ip(eth_name):
    try:
        sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        addr = fcntl.ioctl(sock.fileno(), 0x8915, struct.pack("256s", eth_name))
        return socket.inet_ntoa(addr[20:24])
    except Exception as e:
        logging.error(f"get localhost ip address error: {str(e)}")
        return "127.0.0.1"


def pull_history(http_get_params=""):
    http_client = ""
    try:
        http_client = http.client.HTTPConnection(OMSERVER_ADDRESS, 80, timeout=CONNECTION_TIMEOUT)
        http_client.request("GET", http_get_params)
        response = http_client.getresponse()

        if response.status != 200:
            logging.error(f"response http status error: {response.status}")
            sys.exit()

        http_content = response.read().strip()
        if http_content != "OK":
            logging.error(f"response http content error: {http_content}")
            sys.exit()

    except Exception as e:
        logging.error(f"pull history error: {str(e)}")
        sys.exit()

    finally:
        if http_client:
            http_client.close()
        else:
            logging.error("close http unknown error: {str(}")
            sys.exit()


local_ip = get_local_ip(NET_DRIVER)
user = sys.argv[2]
history_id = sys.argv[1]
history_date = sys.argv[3]
history_time = sys.argv[4]
history_command = ""


for i in range(5, len(sys.argv)):
    history_command += sys.argv[i] + ""

params = "/omaudit/omaudit_pull/?history_id="+history_id+"&history_ip="+local_ip+"&history_user="+user+"&history_datetime="+history_date+urllib.parse.quote(" ")+history_time+"&history_command="+urllib.parse.quote(history_command.strip())

pull_history(params)

