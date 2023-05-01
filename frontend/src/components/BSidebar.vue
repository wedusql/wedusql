<script lang="ts" setup>
import Tree, { TreeNode } from 'primevue/tree';
import { onMounted, ref, watchEffect } from 'vue';

import { Run } from "../../wailsjs/go/queries/Query.js";

const tables = ref([] as TreeNode[]);
const treeSelect = ref({});
const props = defineProps<{
  activeTable: string,
}>();

const emits = defineEmits(['update:activeTable']);

const getAllTables = async () => {
  const qTables = await Run("SHOW TABLES");
  tables.value = [
    {
      key: "0",
      label: "Tables",
      children: qTables.Rows.map((t: any) => {
        return {
          key: `tables-${t[Object.keys(t)[0]]}`,
          label: t[Object.keys(t)[0]],
        };
      }),
    },
  ];
};

onMounted(() => {
  getAllTables();
});

watchEffect(() => {
  const tableKey = Object.keys(treeSelect.value)[0];
  const tableNameFromKey = tableKey?.split("tables-");
  if (!tableNameFromKey || tableNameFromKey.length < 2) {
    return;
  }

  emits('update:activeTable', tableNameFromKey[1]);
});
</script>

<template>
  <Tree
    :value="tables"
    class="w-full h-screen overflow-x-auto"
    selectionMode="single"
    v-model:selectionKeys="treeSelect"
  />
</template>
