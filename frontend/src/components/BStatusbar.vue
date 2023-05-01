<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { GetConnectionInfo } from '../../wailsjs/go/connections/Connection.js';

const props = defineProps<{
  activeTable: string;
}>();

const connectionInfo = ref({} as any);

const getConnectionDsn = async () => {
  const ci = await GetConnectionInfo();
  connectionInfo.value = ci;
};

onMounted(() => {
  getConnectionDsn();
});

</script>

<template>
  <div class="statusbar p-1">
    {{ connectionInfo.User }} @ {{ connectionInfo.Host }} /
    {{ connectionInfo.Database }}
    <span v-if="!!activeTable"> / {{ activeTable }} </span>
  </div>
</template>
