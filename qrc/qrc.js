const args = require('minimist')(process.argv.slice(2))
const QRCode = require("qrcode");
async function main(str) {
    address = str;
    console.log(
        await QRCode.toString(address, { type: "terminal", small: true })
    );
}
main(args['lnu']);