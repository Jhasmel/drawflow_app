import Swal from "sweetalert2";

export default function convertToPython(df, module, prefix = "") {
  let code = "";
  const modules = df.drawflow.drawflow;
  const nodes = modules[module].data;

  if (nodes.length === 0) {
    return code;
  }
  const arr = [];
  Object.keys(nodes).forEach((id) => {
    arr.push(nodes[id]);
  });

  const nodesList = arr.sort((a, b) => {
    return a.pos_x > b.pos_x;
  });

  nodesList.forEach((node) => {
    pythonString(node.name, node, df);
  });

  const nodesFil = nodesList.filter((node) => {
    return (
      node.name === "Operations" ||
      node.name === "Assign" ||
      node.name === "If" ||
      node.name === "Else" ||
      node.name === "For" ||
      node.name === "Print" ||
      node.name === "codeBlock"
    );
  });

  for (let i = nodesFil.length - 1; i >= 0; i--) {
    if (nodesFil.length === 0) {
      break;
    }
    const node = nodesFil.pop();
    let newCode = "";
    removeConnection(nodesFil, node);
    switch (node.name) {
      default:
        newCode = prefix + node.data.codePy;
        break;
    }

    code = newCode + code;
  }
  return code;
}

function removeConnection(nodes, selectedNode) {
  if (nodes.length === 0) {
    return;
  }
  const inputs = selectedNode.inputs;
  Object.keys(inputs).forEach((input) => {
    const connectedNode = nodes.find(
      (node) => node.id === parseInt(inputs[input].connections[0].node)
    );
    if (connectedNode !== undefined) {
      removeConnection(nodes, connectedNode);
      nodes.splice(nodes.indexOf(connectedNode), 1);
    }
  });
}

const pythonString = (name, node, df) => {
  switch (name) {
    case "Number":
      numberPython(node, df);
      break;
    case "Operations":
      operationPython(node, df);
      break;
    case "Assign":
      assignPython(node, df);
      break;
    case "If":
      conditionalPython(node, df);
      break;
    case "Else":
      elsePython(node, df);
      break;
    case "For":
      forPython(node, df);
      break;
    case "Print":
      printPython(node, df);
      break;
    case "codeBlock":
      codeBlockPython(node, df);
      break;
  }
};

function numberPython(node, df) {
  if (isNaN(node.data.result) || node.data.result === null) {
    node.data.result = node.data.codePy = 0;
    df.updateNodeDataFromId(node.id, node.data);
  } 
}

function operationPython(node, df) {
  let result = 0;
  const data = node.data;
  const number = [];
  const expressions = [];
  const inputs = node.inputs;

  try {
    Object.keys(inputs).forEach((key) => {
      const input = inputs[key];
      const connection = input.connections[0];

      const currentNode = df.getNodeFromId(connection.node);

      number.push(currentNode.data.result);
      expressions.push(currentNode.data.codePy);
    });
  } catch {
    Swal.fire('Connect the OperationsN to the NumberN.')
  }

  switch (data.method) {
    case "+":
      result = number[0] + number[1];
      break;
    case "-":
      result = number[0] - number[1];
      break;
    case "*":
      result = number[0] * number[1];
      break;
    case "/":
      result = number[0] / number[1];
      break;
  }

  data.result = result;
  data.codePy = `${expressions[0]} ${data.method} ${expressions[1]}`;
  df.updateNodeDataFromId(node.id, data);
}

function assignPython(node, df) {
  const data = node.data;
  if (!isNaN(data.assign) || data.assign === undefined) {
    Swal.fire('AssigN is undefined. Please connect to another node and do not assign a number')
    return;
  }

  try {
    const connectedNode = df.getNodeFromId(
      node.inputs.input_1.connections[0].node
    );

    data.value = connectedNode.data.result;
    data.codePy = `${data.assign} = ${connectedNode.data.codePy}\n`;

    df.updateNodeDataFromId(node.id, data);
  } catch {
    Swal.fire('AssigN is undefined. Please connect to another node and do not assign a number')
  }
}

function conditionalPython(node, df) {
  let result = 0;
  const data = node.data;
  const number = [];
  const expressions = [];
  const inputs = node.inputs;

  try {
    Object.keys(inputs).forEach((key) => {
      const input = inputs[key];
      const connection = input.connections[0];

      const currentNode = df.getNodeFromId(connection.node);

      number.push(currentNode.data.result);
      expressions.push(currentNode.data.codePy);
    });
  } catch {
    Swal.fire('It must be connected to the number node')
  }

  switch (data.method) {
    case ">":
      result = number[0] > number[1];
      break;
    case "<":
      result = number[0] < number[1];
      break;
    case ">":
      result = number[0] >= number[1];
      break;
    case "<":
      result = number[0] <= number[1];
      break;
    case "==":
      result = number[0] == number[1];
      break;
    case "!=":
      result = number[0] != number[1];
      break;
  }

  data.result = result;
  data.codePy = `if ${expressions[0]} ${data.conditional} ${expressions[1]}: print(${node.data.conditionalNumber}) \n`;
  df.updateNodeDataFromId(node.id, data);
}

function elsePython(node, df) {
  const data = node.data;
  const connection = node.inputs.input_1.connections[0];
  if (connection) {
    const connectedNode = df.getNodeFromId(connection.node);

    data.codePy =
      connectedNode.data.codePy + `else: print(${node.data.result})\n`;
  } else {
    Swal.fire('Please, connect to another to another node')
  }
  df.updateNodeDataFromId(node.id, data);
}

function forPython(node, df) {
  let result = 0;
  const data = node.data;
  const number = [];
  const expressions = [];
  const inputs = node.inputs;

  try {
    Object.keys(inputs).forEach((key) => {
      const input = inputs[key];
      const connection = input.connections[0];

      const currentNode = df.getNodeFromId(connection.node);

      number.push(currentNode.data.result);
      expressions.push(currentNode.data.codePy);
    });
  } catch {
    Swal.fire('It must be connected to the number node')
  }

  data.result = result;
  data.codePy = `for i in range (${expressions[0]}, ${expressions[1]}): print(${node.data.printResult}) \n`;
  df.updateNodeDataFromId(node.id, data);
}

function printPython(node, df) {
  const data = node.data;
  const connection = node.inputs.input_1.connections[0];
  if (connection) {
    const connectedNode = df.getNodeFromId(connection.node);

    data.codePy = `print(${connectedNode.data.codePy})\n`;
  } else {
    Swal.fire('Connect to another node')
    data.codePy = `print()\n`;
  }
  df.updateNodeDataFromId(node.id, data);
}

function codeBlockPython(node, df) {
  const data = node.data;

  data.codePy = node.data.writtenCode;
  df.updateNodeDataFromId(node.id, data);
}
