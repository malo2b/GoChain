<template>
  <div>
    <Error v-if="state.error" :error="state.error"/>
    <div>

      <div v-if="errors.length">
        <b>Veuillez réessayer</b>
        <ul>
          <li v-for="error in errors" :key="error.key">{{ error }}</li>
        </ul>
      </div>

      <p>
        <label for="from">Envoyeur</label>
        <input placeholder="De" id="from" v-model="form.from" type="text" name="from">
      </p>

      <p>
        <label for="to">Receveur</label>
        <input placeholder="À" id="to" v-model="form.to" type="text" name="to">
      </p>

      <p>
        <label for="value">Valeur</label>
        <input placeholder="Montant" id="value" v-model="form.value" type="text" name="value">
      </p>

      <p>
        <button class="bg-blue-400" v-on:click="checkForm"> Créer</button>
      </p>

    </div>

  </div>
</template>

<script>
import axios from "@/api/axios";
import Error from "@/components/Error";


export default {
  name: "TxCreateView",
  components:{
    Error
  },
  data() {
    return {
      state: {
        error: null,
        success: false,
        loading: true
      },
      form: {
        from: null,
        to: null,
        value: null
      },
      response: null,
      errors: []
    }
  },
  methods: {
    addTx() {
      var bodyFormData = new FormData();
      bodyFormData.append('from', this.form.from);
      bodyFormData.append('to', this.form.to);
      bodyFormData.append('value', this.form.value);

      axios({
        method: "post",
        url: 'addTx',
        data: bodyFormData,
        headers: {"Content-Type": "multipart/form-data"},
      }).then((res) => {
        console.log(res)
        //this.reset();
        //this.state.success = true;
      }).catch((error) => {
        console.log(error.response)
        this.state.error = error.response.data;
      }).finally(() => {
        //this.state.loading =  false;
      });
    },
    checkForm() {
      this.errors = [];

      if (!this.form.from) {
        this.errors.push("Le nom du compte de départ est requis.");
      }
      if (!this.form.to) {
        this.errors.push('Le nom du compte d\'arrivé est requis.');
      }
      if (!this.form.value) {
        this.errors.push('La valeur de la transaction est requise.');
      }

      if (!this.errors.length) {
        this.addTx()
      }

    },
  }
}
</script>

<style scoped>

</style>