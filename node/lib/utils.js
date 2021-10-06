const client = require("../connection/client");

function insert(index, type, data) {
  return client.index({
    index,
    type,
    body: data,
  });
}
exports.insert = insert;

function search(index, query) {
  return client.search({
    index,
    body: {
      query: {
        match: { body: query.keyword },
      },
    },
  });
}

exports.search = search;
