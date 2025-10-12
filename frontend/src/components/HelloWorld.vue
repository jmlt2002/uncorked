<template>
  <div>
    <IftaLabel>
      <InputText id="usernameEmail" v-model="usernameEmail" variant="filled" />
      <label for="usernameEmail">Email or username</label>
    </IftaLabel>

    <IftaLabel>
      <Password id="password" v-model="password" variant="filled" />
      <label for="password">Password</label>
    </IftaLabel>

    <Button label="Submit" class="btn"/>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { InputText, Password, IftaLabel, Button } from 'primevue';

const usernameEmail = ref('');
const password = ref('');
</script>

<style scoped>
.btn {
  width: 100%;
}
</style>
