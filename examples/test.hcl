job "test2" {
  datacenters = ["dc1"]
  type        = "service"

  group "test" {
    restart {
      attempts = 0
      mode     = "fail"
    }
    task "test01" {
      driver = "firecracker-task-driver"
      config {
       KernelImage = "/home/neirac/rootfs/hello-vmlinux.bin" 
       Firecracker = "/home/neirac/versions/firecracker" 
       Vcpus = 1 
       Mem = 128
       Network = "default"
       BootDisk = "/home/neirac/rootfs/hello-rootfs.ext4"
      }
    }
  }
}
