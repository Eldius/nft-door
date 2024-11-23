
process.exit(0).catch((error) => { console.error(error); process.exit(1); });

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Deploying contracts with the account:", deployer.address);
    const MyContract = await ethers.getContractFactory("DoorControl");
    const myContract = await MyContract.deploy();
    console.log("MyContract deployed to:", myContract.address);
}
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
