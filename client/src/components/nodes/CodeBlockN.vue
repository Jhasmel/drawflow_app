<template>
  <div ref="el">
    <nodeHeader title="codeBlock" />
    <el-input 
    @change="updateInput"
    v-model="writtenCode"
    df-writtenCode
    placeholder="insert code"
    />
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
    const writtenCode = ref("");
    const dataNode = ref({});

    let df = getCurrentInstance().appContext.config.globalProperties.$df.value;

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);

      writtenCode.value = dataNode.value.data.writtenCode;

      })

    const updateInput = () => {
      dataNode.value.data.writtenCode = writtenCode.value;
      dataNode.value.data.codePy = writtenCode.value;
      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    }

    return {
      el,
      updateInput,
      nodeId,
      writtenCode,
    };
  },
});
</script>
