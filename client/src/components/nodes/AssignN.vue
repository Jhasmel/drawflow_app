<template>
  <div ref="el">
    <nodeHeader title="Assign" />
    <el-input
      v-model="assignVariable"
      placeholder="Assign the variable"
      df-assign
      type="text"
      @change="updateInput"
    />
    <div>{{ assignValue }}</div>
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
    const dataNode = ref({});
    const assignVariable = ref("");
    const assignValue = ref(0);
    const internalInstance = getCurrentInstance();

    let df = internalInstance.appContext.config.globalProperties.$df.value;

    const updateInput = () => {
      dataNode.value.data.assign = assignVariable.value;
      dataNode.value.data.value = assignValue.value;

      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    };
    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);
      assignVariable.value = dataNode.value.data.assign;
      assignValue.value = dataNode.value.data.value;
    });

    return {
      el,
      nodeId,
      updateInput,
      assignVariable,
      assignValue,
    };
  },
});
</script>
