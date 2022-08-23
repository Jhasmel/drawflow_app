<template>
  <div ref="el">
    <nodeHeader title="codeBlock" />
    <el-button @click="blockCodeButton">Enter Code</el-button>
  </div>
</template>

<script>
import {
  defineComponent,
  nextTick,
  onMounted,
  ref,
  getCurrentInstance,
} from "vue";
import nodeHeader from "../NodeHeaderN.vue"

export default defineComponent({
  components: { nodeHeader },
  setup() {
    const el = ref(null);
    const nodeId = ref(0);
    const dataNode = ref({});

    let df = getCurrentInstance().appContext.config.globalProperties.$df.value;

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);

      if (df.drawflow.drawflow[`code-block-${nodeId.value}`] === undefined) {
        df.addModule(`code-block-${nodeId.value}`);
      }
    });
    const blockCodeButton = () => {
      df.changeModule(`code-block-${nodeId.value}`);
    }

    return {
      el,
      blockCodeButton,
      nodeId,
    };
  },
});
</script>
