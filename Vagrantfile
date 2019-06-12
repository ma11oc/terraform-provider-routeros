Vagrant.configure("2") do |config|
  config.vm.box = "cheretbe/routeros"
  config.vm.box_version = "6.40.6"
  config.vm.network "forwarded_port", guest: 8728, host: 8728
  config.vm.network "forwarded_port", guest: 8729, host: 8729
end
