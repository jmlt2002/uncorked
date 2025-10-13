<template>
  <div class="container">
    <Form class="form">
      <div class="form-field">
        <InputText id="usernameEmail" v-model="usernameEmail" placeholder="Username or email" />
      </div>

      <div class="form-field">
        <Password id="password" v-model="password" placeholder="Password" fluid/>
      </div>

      <Button label="Submit" class="form-field" />

      <p>Not a member yet?<a href=""> Register here!</a></p>
    </Form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { InputText, Password, Button } from 'primevue';
import { Form } from '@primevue/forms';

const usernameEmail = ref('');
const password = ref('');
</script>

<style scoped>
.container {
  min-height: 100vh;
  background-image: url('@/assets/cellar_blurred.jpg');
  background-size: cover;
  background-position: center;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100vw;
}

.form {
  background-color: rgba(255, 255, 255, 0.9);
  padding: 2rem;
  border-radius: 10px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
  width: 50%;
  box-sizing: border-box;
  max-width: 400px;

  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 1em;
}

.form-field {
  display: flex;
  flex-direction: column;
  width: 90%;
}

</style>
