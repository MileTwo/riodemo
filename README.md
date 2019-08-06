# riodemo
Demo rio 



```bash
rio run -p 8080/http --name rdemo monachus/rancher-demo:yellow
rio scale default/rdemo=1-5
rio scale default/rdemo=3
rio stage --image=monachus/rancher-demo:pink default/rdemo:v1
rio weight --rollout-interval 2 default/rdemo:v1=25
rio promote --rollout-interval 2 default/rdemo:v1
```