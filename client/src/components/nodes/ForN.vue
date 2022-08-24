<template>
  <div ref="el">
    <nodeHeader title="For" />
    <el-input 
    @change="updateInput"
    v-model="printResult"
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
    const printResult = ref(0);
    const dataNode = ref({});

    let df = getCurrentInstance().appContext.config.globalProperties.$df.value;

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);

      printResult.value = dataNode.value.data.printResult;
    });

    const updateSelect = () =>{
      dataNode.value.data.conditional = conditional.value;
      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    }

    const updateInput = () => {
      dataNode.value.data.printResult = printResult.value;
      dataNode.value.data.codePy = `${printResult.value}`;

      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    }

    return {
      el,
      nodeId,
      printResult,
      updateSelect,
      updateInput,
    };
  },
});
</script>
