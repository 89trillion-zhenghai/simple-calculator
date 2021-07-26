from locust import HttpUser, TaskSet, task,between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def baidu_index1(self):
       self.client.post("/calculator",data={'str':'2+3+7*1-7/2+12*2'})