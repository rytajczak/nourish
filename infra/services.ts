const vpc = new sst.aws.Vpc("vpc", { bastion: true });
const cluster = new sst.aws.Cluster("services", { vpc });
