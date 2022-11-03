make uninstall
make install
kubectl apply -f config/samples/lb_v1_nic.yaml
kubectl apply -f config/samples/lb_v1_virnet.yaml
kubectl apply -f config/samples/lb_v1_virtualserver.yaml
kubectl apply -f config/samples/lb_v1_vm.yaml
make run