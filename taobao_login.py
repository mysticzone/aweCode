"""
-------------------------------------------------
   FileName：     taobao_login
   Date：         2022/12/7
-------------------------------------------------
"""

from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC


class TaobaoInfo(object):
    def __init__(self):
        self.url = "https://login.taobao.com/member/login.html"
        options = webdriver.ChromeOptions()
        options.add_experimental_option("prefs", {"profile.managed_default_content_settings.images": 2})
        options.add_experimental_option("excludeSwitchs", ["enable-automation"])

        self.browser = webdriver.Chrome(executable_path=chrome_webdriver, options=options)
        # Set timeout
        self.wait = WebDriverWait(self.browser, 10)


    def login(self):
        self.browser.get(self.url)

        pass_login = self.wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, ".qrcode-login > .login-links > .forget-pwd")))
        pass_login.click()

        weibo_login = self.wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, ".weibo-login")))
        weibo_login.click()

        weibo_user = self.wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, ".username > .W_input")))
        weibo_user.send_keys(weibo_username)

        weibo_pwd = self.wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, ".password > .W_input")))
        weibo_pwd.send_keys(weibo_password)

        submit = self.wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, ".btn_tip > a > span")))
        submit.click()

        taobao_name = self.wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, ".site-nav-bd > .ul.site-nav-bd-l > li#J_SiteNavLogin > div.site-nav-menu-hd > div.site-nav-user > a.site-nav-login-info-nick")))
        print(taobao_name.text)

if __name__ == "__main__":
    chrome_webdriver = "./chromedriver"
    weibo_username = "微博账号"
    weibo_password = "微博密码"

    tb_info = TaobaoInfo()
    tb_info.login()
