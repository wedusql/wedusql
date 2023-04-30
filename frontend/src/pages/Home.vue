<script lang="ts" setup>
import { onMounted, reactive, ref } from "vue";
import Tree, { TreeNode } from "primevue/tree";

import { Run } from '../../wailsjs/go/queries/Query.js';

const tables = ref([] as TreeNode[]);

const getAllTables = async () => {
  const qTables = await Run('SHOW TABLES');
  tables.value = [
    {
      key: "0",
      label: "Tables",
      children: qTables.Rows.map((t: any, i: any) => {
        return {
          key: `tables-${i}`,
          label: t.Tables_in_mysql,
        };
      }),
    },
  ];
  console.log(qTables.Rows);
}

onMounted(() => {
  getAllTables();
});

</script>

<template>
  <div class="flex">
    <div class="w-3">
      <Tree :value="tables" class="w-full h-screen overflow-x-auto" />
    </div>

    <div>asdasd</div>
  </div>
</template>
