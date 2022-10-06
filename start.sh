# First start consul
#consul agent -dev -bind 0.0.0.0 -client 0.0.0.0  &
# Now nomad agent
nomad agent -dev -config=/home/neirac/firecracker-task-driver/config.hcl -data-dir=/home/neirac/firecracker-task-driver -plugin-dir=/home/neirac/firecracker-task-driver/plugin -bind=0.0.0.0
