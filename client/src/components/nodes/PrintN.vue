<template>
  <div ref="el">
    <nodeHeader title="Print" />
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

    let df = getCurrentInstance().appContext.config.globalProperties.$df.value;

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);
    });

    return {
      el,
      nodeId,
    };
  },
});
</script>
