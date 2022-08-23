<template>
  <div ref="el">
    <nodeHeader title="Operations" />
    <el-select
      v-model="operations"
      @change="updateSelect"
      placeholder="Select"
      df-method
    >
      <el-option value="+">Sum</el-option>
      <el-option value="-">Rest</el-option>
      <el-option value="*">Multiply</el-option>
      <el-option value="/">Divide</el-option>
    </el-select>
    <span df-result> {{ result }}</span>
  </div>
</template>

<script>
import {
  defineComponent,
  onMounted,
  getCurrentInstance,
  ref,
  nextTick,
} from "vue";
import nodeHeader from "../NodeHeaderN.vue"

export default defineComponent({
  components: {
    nodeHeader,
  },
  setup() {
    const el = ref(null);
    const nodeId = ref(0);
    const operations = ref("-");
    const dataNode = ref({});
    let result = ref(0);

    let df = getCurrentInstance().appContext.config.globalProperties.$df.value;

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);

      result.value = dataNode.value.data.result;
      operations.value = dataNode.value.data.method;
    });

    const updateSelect = () => {
      dataNode.value.data.method = operations.value;
      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    }

    return {
      el,
      operations,
      result,
      nodeId,
      updateSelect,
    };
  },
});
</script>

<style></style>
