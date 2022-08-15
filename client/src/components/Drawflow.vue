<template>
  <el-container>
    <el-header class="header">
      <h3>{{ currentModule }}</h3>
      <el-button type="primary" @click="clearProgram">Clear</el-button>
    </el-header>


    <el-container class="program">
      <el-button @click="saveProgram">Save Program</el-button>
      <el-button @click="handleclickList">List Programs</el-button>
      <el-button type="primary" @click="exportPython"
        >Generate Python Code</el-button
      >
      <el-button
        v-if="currentModule !== 'Home'"
        type="success"
        @click="returnHomeModule"
      >
        Return Home</el-button
      >
    </el-container>

    <el-container class="container">
      <el-aside width="200px" class="column">
        <ul>
          <li
            v-for="n in listNodes"
            :key="n"
            draggable="true"
            :data-node="n.item"
            @dragstart="drag($event)"
            class="drag-drawflow"
          >
            <div class="node" :style="`background: #6290C8`">
              {{ n.item }}
            </div>
          </li>
        </ul>
      </el-aside>
      <el-main>
        <div
          id="drawflow"
          @drop="drop($event)"
          @dragover="allowDrop($event)"
        ></div>
      </el-main>
      <div>
        <el-drawer v-model="listPrograms" title="List of Programs">
  <ul>
    <li
      class="item"
      v-for="program in totalPrograms"
      :key="program.uid"
    >
      <div>{{ program.name }}</div>
      <div>
          <div>
            <el-button
              @click="deleteProgram(program)"
              type="danger"
              >Delete</el-button
            >
            <el-button
              @click="obtainProgram(program.uid)"
              type="success"
              >Obtain</el-button 
            >
          </div>
      </div>
    </li>
  </ul>
        </el-drawer>
      </div>
    </el-container>
  </el-container>

  <el-dialog v-model="codePython" title="Python Code">
    <el-container>
      <div>
        <runCode :code="PythonCode" />
        <el-button @click="codePython = false">Cancel</el-button>
      </div>
    </el-container>
  </el-dialog>
</template>

<script>
import Drawflow from "drawflow";
import styleDrawflow from "drawflow/dist/drawflow.min.css";
import {
  onMounted,
  shallowRef,
  h,
  getCurrentInstance,
  render,
  readonly,
  ref,
} from "vue";
import Number from "./nodes/NumberN.vue";
import Operations from "./nodes/OperationsN.vue";
import Assign from "./nodes/AssignN.vue";
import If from "./nodes/IfN.vue";
import Else from "./nodes/ElseN.vue";
import For from "./nodes/ForN.vue";
import Print from "./nodes/PrintN.vue";
import codeBlock from "./nodes/CodeBlockN.vue";
import pythonString from "./functions/convert_to_python";
import Swal from "sweetalert2";
import runCode from "./RunCode.vue";

export default {
  name: "drawflow",
  setup() {
    const listNodes = readonly([
      {
        name: "Number",
        item: "Number",
        input: 0,
        output: 1,
      },
      {
        name: "Operations",
        item: "Operations",
        input: 2,
        output: 1,
      },
      {
        name: "Assign",
        item: "Assign",
        input: 1,
        output: 0,
      },
      {
        name: "If",
        item: "If",
        input: 2,
        output: 1,
      },
      {
        name: "Else",
        item: "Else",
        input: 1,
        output: 0,
      },
      {
        name: "For",
        item: "For",
        input: 0,
        output: 1,
      },
      {
        name: "Print",
        item: "Print",
        input: 1,
        output: 0,
      },
      {
        name: "codeBlock",
        item: "codeBlock",
        input: 1,
        output: 0,
      },
    ]);
    const totalPrograms = ref([]);
    const editor = shallowRef({});
    const codePython = ref(false);
    const PythonCode = ref("");
    const currentModule = ref("Home");
    const listPrograms = ref(false);
    const Vue = { version: 3, h, render };
    const internalInstance = getCurrentInstance();

    internalInstance.appContext.app._context.config.globalProperties.$df =
      editor;

    function obtainProgram (uid) {
            // Swal.fire(
      //   'Obtained!',
      //   'Your file has been obtained.',
      //   'success')
        // .then(() => {
      fetch(`http://localhost:9000/v1/programs/${uid}`)
      .then((response) => response.json())
      .then((query) => {
        if (query.getById.length !== 0) {
          const programSelected = query.getById[0]
          const exportCode = programSelected.code
         console.log(exportCode)

        }
        // if (query.error_code) {
        //   Swal.fire('It could not connect to the database.')
        //   throw new Error("It could not connect to the database.");
        // }
          // console.log(query);
      })

      // .then((queryID) => {
      //   if (queryID.getById.length != 0) {
      //     const programSelected = queryID.getById[0]
      //     const drawflow = JSON.parse(programSelected.code)
      //     editor.value.import(drawflow) 
      //   }
      // })
    }

    function deleteProgram (codeDeleted) {
      Swal.fire({
        title: 'Are you sure?',
        text: "You won't be able to revert this!",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Yes, delete it!'
        }).
        then((result) => {
          if (result.isConfirmed) {
            Swal.fire(
              'Deleted!',
              'Your file has been deleted.',
              'success'
            )
          }
        }).
        then(() => {
          fetch("http://localhost:9000/v1/programs", {
            method: "DELETE",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(codeDeleted),
          })
        })
      }

    function exportPython() {
      PythonCode.value = pythonString(editor.value, "Home");
      codePython.value = true;
      editor.value.changeModule("Home");
    }

    function clearProgram() {
      editor.value.import({ drawflow: { Home: { data: {} } } });
    }

    function returnHomeModule() {
      editor.value.changeModule("Home");
    }
    
    function saveProgram() {
      const exportprogram = pythonString(editor.value, "Home");
        const datapayload = {
          uid: "_:program",
          name: "",
          code: exportprogram,
        };
         Swal.fire({
        title: "Name of the Code",
        input: "text",
        inputLabel: "Name",
        showCancelButton: true,
        inputValidator: (value) => {
          if (!value) {
            return "Please, submit a name!";
          }
        },
      }).then(({ value }) => {
          datapayload.name = value;
          if (totalPrograms.value.some((program) => program.name === value)) {
            Swal.fire('The name is already registered!')
            return;
          }
            fetch("http://localhost:9000/v1/programs/", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(datapayload),
          })
            .then((response) => response.json())
            .then((query) => {
              Swal.fire(
                'Code saved!',
                '',
                'success'
              ),   
              datapayload.uid = query.uids.program;
              const newPrograms = totalPrograms.value;
              newPrograms.push(datapayload);
              totalPrograms.value = newPrograms;
            })
            .catch((err) => {
              console.log(err);
              Swal.fire('Program could not be saved.')
            });
        });
        return;
        }   

    function handleclickList () {
      listPrograms.value = true;
      editor.value.changeModule("Home");


    }
    const drag = (ev) => {
      if (ev.type === "touchstart") {
        mobile_item_selec = ev.target
          .closest(".drag-drawflow")
          .getAttribute("data-node");
      } else {
        ev.dataTransfer.setData("node", ev.target.getAttribute("data-node"));
      }
    };
    const drop = (ev) => {
      if (ev.type === "touchend") {
        var parentdrawflow = document
          .elementFromPoint(
            mobile_last_move.touches[0].clientX,
            mobile_last_move.touches[0].clientY
          )
          .closest("#drawflow");
        if (parentdrawflow != null) {
          addNodeToDrawFlow(
            mobile_item_selec,
            mobile_last_move.touches[0].clientX,
            mobile_last_move.touches[0].clientY
          );
        }
        mobile_item_selec = "";
      } else {
        ev.preventDefault();
        var data = ev.dataTransfer.getData("node");
        addNodeToDrawFlow(data, ev.clientX, ev.clientY);
      }
    };
    const allowDrop = (ev) => {
      ev.preventDefault();
    };
    let mobile_item_selec = "";
    let mobile_last_move = null;
    function positionMobile(ev) {
      mobile_last_move = ev;
    }
    function addNodeToDrawFlow(name, pos_x, pos_y) {
      pos_x =
        pos_x *
          (editor.value.precanvas.clientWidth /
            (editor.value.precanvas.clientWidth * editor.value.zoom)) -
        editor.value.precanvas.getBoundingClientRect().x *
          (editor.value.precanvas.clientWidth /
            (editor.value.precanvas.clientWidth * editor.value.zoom));
      pos_y =
        pos_y *
          (editor.value.precanvas.clientHeight /
            (editor.value.precanvas.clientHeight * editor.value.zoom)) -
        editor.value.precanvas.getBoundingClientRect().y *
          (editor.value.precanvas.clientHeight /
            (editor.value.precanvas.clientHeight * editor.value.zoom));
      const nodeSelected = listNodes.find((ele) => ele.item == name);
      editor.value.addNode(
        name,
        nodeSelected.input,
        nodeSelected.output,
        pos_x,
        pos_y,
        name,
        {},
        name,
        "vue"
      );
    }
    onMounted(() => {
      var elements = document.getElementsByClassName("drag-drawflow");
      for (var i = 0; i < elements.length; i++) {
        elements[i].addEventListener("touchend", drop, false);
        elements[i].addEventListener("touchmove", positionMobile, false);
        elements[i].addEventListener("touchstart", drag, false);
      }
      const id = document.getElementById("drawflow");
      editor.value = new Drawflow(
        id,
        Vue,
        internalInstance.appContext.app._context
      );
      editor.value.start();
      editor.value.registerNode("Number", Number, {}, {});
      editor.value.registerNode("Operations", Operations, {}, {});
      editor.value.registerNode("Assign", Assign, {}, {});
      editor.value.registerNode("If", If, {}, {});
      editor.value.registerNode("Else", Else, {}, {});
      editor.value.registerNode("For", For, {}, {});
      editor.value.registerNode("Print", Print, {}, {});
      editor.value.registerNode("codeBlock", codeBlock, {}, {});

      editor.value.on("connectionCreated", (conn) => {
        const nodeOutput = editor.value.getNodeFromId(conn.output_id);
        const nodeInput = editor.value.getNodeFromId(conn.input_id);
        if (nodeOutput.outputs[conn.output_class].connections.length > 1) {
          const removeConnection =
            nodeOutput.outputs[conn.output_class].connections[0];
          editor.value.removeSingleConnection(
            conn.output_id,
            removeConnection.node,
            conn.output_class,
            removeConnection.output
          );
        } else if (nodeInput.inputs[conn.input_class].connections.length > 1) {
          const removeConnection =
            nodeInput.inputs[conn.input_class].connections[0];
          editor.value.removeSingleConnection(
            removeConnection.node,
            conn.input_id,
            removeConnection.input,
            conn.input_class
          );
        }
      });
      editor.value.on("moduleChanged", (name) => {
        currentModule.value = name;
      });
      fetch("http://localhost:9000/v1/programs")
        .then((response) => response.json())
        .then((query) => {
          if (query.error_code) {
            Swal.fire('It could not connect to the database.')
            throw new Error("It could not connect to the database.");
          }
          console.log(query);
          const tprograms = query.programs;
          totalPrograms.value = tprograms;
        })
        .catch((err) => {});
    });
    return {
      exportPython,
      listNodes,
      drag,
      drop,
      allowDrop,
      codePython,
      PythonCode,
      editor,
      saveProgram,
      totalPrograms,
      currentModule,
      returnHomeModule,
      clearProgram,
      handleclickList,
      deleteProgram,
      obtainProgram,
      listPrograms,
    };
  },
  components: { runCode },
};
</script>
<style scoped>
.h3 {
  color: black;
}
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid white;
}

.program {
  display: flex;
  justify-content: center;
  align-items: center;
  border-bottom: 1px solid white;
  height: 60px;
}

.item {
  color: black;
}
.buttons {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.container {
  min-height: calc(100vh - 100px);
}
.column {
  border-right: 1px solid;
}
.column ul {
  padding-inline-start: 0px;
  padding: 10px 10px;
}
.column li {
  background: transparent;
}
.node {
  border: 1px solid white;
  display: block;
  height: 55px;
  line-height: 35px;
  padding: 10px;
  margin: 10px 0px;
  cursor: move;
  color: black;
  align-items: center;
}
#drawflow {
  width: 100%;
  height: 100%;
  text-align: initial;
  background: #2b2c30;
  background-size: 20px 20px;
  background-image: radial-gradient(#494949 1px, transparent 1px);
}
</style>
