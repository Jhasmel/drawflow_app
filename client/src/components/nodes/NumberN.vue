<template>
  <div ref="el">
    <nodeHeader title="Number" />
    <el-input
      @change="updateInput"
      v-model="result"
      placeholder="Pick a number"
    />
  </div>
</template>

<script>
import {
  defineComponent,
  getCurrentInstance,
  nextTick,
  onMounted,
  ref,
} from "vue";
import nodeHeader from "../NodeHeaderN.vue"

export default defineComponent({
  components: {
    nodeHeader,
  },
  setup() {
    const el = ref(null);
    const nodeId = ref(0);
    const result = ref(0);
    const dataNode = ref({});

    let df = getCurrentInstance().appContext.config.globalProperties.$df.value;

    onMounted(async () => {
      await nextTick();

      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);
      result.value = dataNode.value.data.result;
    });

    const updateInput = () => {
      dataNode.value.data.result = parseInt(result.value);
      dataNode.value.data.codePy = `${result.value}`;

      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    };

    return {
      result,
      el,
      updateInput,
      nodeId,
    };
  },
});
</script>
