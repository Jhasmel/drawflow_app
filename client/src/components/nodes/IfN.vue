<template>
  <div ref="el">
    <nodeHeader title="If" />
    <el-select
      v-model="conditional"
      @change="updateSelect"
      placeholder="Select"
      df-conditional
    >
      <el-option value=">">Greater than</el-option>
      <el-option value="<">Lesser than</el-option>
      <el-option value=">=">Greater than or equal to</el-option>
      <el-option value="<=">Lesser than or equal to</el-option>
      <el-option value="==">Equal to</el-option>
      <el-option value="!=">Not equal to</el-option>
    </el-select>
    <br />
    <br />
    <el-input     
    @change="updateInput"
    v-model="conditionalNumber"
    placeholder="print"
 />
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
    const conditional = ref("");
    const dataNode = ref({});
    const conditionalNumber = ref(0);


    let df = getCurrentInstance().appContext.config.globalProperties.$df.value;

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);

      conditionalNumber.value = dataNode.value.data.conditionalNumber;
      conditional.value = dataNode.value.data.conditional;
    });

    function updateSelect() {
      dataNode.value.data.conditional = conditional.value;
      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    };

    function updateInput () {
      dataNode.value.data.conditionalNumber = parseInt(conditionalNumber.value);
      dataNode.value.data.codePy = `${conditionalNumber.value}`;

      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    };


    return {
      el,
      nodeId,
      conditional,
      conditionalNumber,
      updateSelect,
      updateInput,
    };
  },
});
</script>

<style></style>
