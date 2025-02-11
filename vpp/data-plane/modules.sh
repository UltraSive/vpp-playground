sudo modprobe uio
sudo modprobe uio_pci_generic
sudo dpdk-devbind.py --bind=uio_pci_generic 00:13.0
dpdk-devbind.py --status