function process(params) {
    params['grade']['int'] = params['grade']['int'] + 1
    return params
}

const definitions = {
    name: 'Student',
    type: 'record',
    fields: [
        {name: 'name', type: ["null", "string"]},
        {name: 'age', type: ["null", "int"]},
        {name: 'grade', type: ["null", "int"]}
    ]
}

module.exports.definitions = definitions