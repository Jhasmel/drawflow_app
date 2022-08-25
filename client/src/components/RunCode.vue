<template>
  <div>
    <code> {{ code }}</code>
    <div>
      <el-button @click="runCode" type="primary">Run Code </el-button>
    </div>
    <br />
    <code>
      <p>{{ codeResult }}</p>
    </code>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";

export default defineComponent({
  props: {
    code: {
      type: String,
    },              
  },
  setup(props) {
    const codeResult = ref("");
                      
    function runCode() {
      fetch("http://localhost:9000/v1/programs/runCode", {
        method: "POST",
        body: props.code,
      })
        .then((response) => response.text())
        .then((text) => {
          if (text.startsWith("Traceback")) {
            codeResult.value = "Invalid code";
          } else {
            codeResult.value = text;
          }
        });
    }
    return {
      codeResult,
      runCode,
    };
  },
});
</script>