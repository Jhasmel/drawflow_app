<template>
  <div ref="el">
    <nodeHeader title="For" />
    <div>From <el-input v-model="From" df-From></el-input></div>
    <div>To <el-input v-model="To" df-To></el-input></div>
  </div>
</template>

<script>
import {
  defineComponent,
  ref,
  onMounted,
  nextTick,
  getCurrentInstance,
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
    const From = ref(0);
    const To = ref(0);

    let df =
      getCurrentInstance().appContext.config.globalProperties.$df.value;

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);
      From.value = dataNode.value.data.From;
      To.value = dataNode.value.data.To;
    });
    return {
      el,
      nodeId,
      From,
      To,
    };
  },
});
</script>
