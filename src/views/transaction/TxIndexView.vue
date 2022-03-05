<template>
  <div>
    <h1 class="text-2xl">Solde des comptes</h1>
    <div v-if="state.success">
      <div class="border border-gray-600 rounded-xl px-4 py-1 m-2 flex" v-for="tx in response" :key="tx.key">

        <div class="w-1/2 flex">
          <div class="font-bold mr-2">Compte: </div>
          {{ tx['account'] }}
        </div>

        <div class="font-bold mr-2">Solde: </div>
        {{ tx['balance'] }}
      </div>
    </div>
    <router-link to="/transaction/create"><button class="bg-blue-500">créer</button>  </router-link>

  </div>
</template>

<script>
import axios from "@/api/axios";

export default {
  name: "TxIndexView",
  data(){
    return{
      state:{
        error: null,
        success: false,
        loading: true
      },
      response: null
    }
  },

  mounted() {
    this.balanceList();
  },
  methods:{
    balanceList(){
      axios.get('balancesList').then((res) => {
        this.response = res.data;
        this.reset();
        this.state.success = true;

      }).catch((error) => {
        this.state.error = error;
        console.log(error)
      }).finally(() => {
        this.state.loading =  false;
      });
    },
    reset() {
      this.state.success = false;
      this.state.error = false;
      this.state.loading = true;
    },

  }
}
</script>

<style scoped>

</style>