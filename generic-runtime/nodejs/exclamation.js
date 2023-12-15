async function process(message, context) {
    const logger = context.getLogger();
    for (let word of message.split(" ")) {
        await context.incrementCounter(word, 1)
        let count = await context.getCounter(word)
        logger.info(`got word: ${word} for ${count['value']} times`)
    }
    await context.publish("persistent://public/default/test-node-package-serde-extra", message.concat("!!!"))
    return message.concat("!");
}