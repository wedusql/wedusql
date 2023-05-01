<script lang="ts" setup>
import Toolbar from 'primevue/toolbar';
import Button from 'primevue/button';
import { useRouter } from 'vue-router';

import { Disconnect } from "../../wailsjs/go/connections/Connection.js";

const router = useRouter();

const disconnect = async() => {
  await Disconnect();
  router.push('/');
};
</script>

<template>
  <Toolbar>
    <template #start>
    </template>

    <template #end>
      <Button label="Disconnect" severity="danger" @click="disconnect()" />
    </template>
  </Toolbar>
</template>
