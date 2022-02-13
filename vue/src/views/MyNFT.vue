<template>
  <div>
    <Balances />
    <Delegations />
    <Undelegations />
  </div>
</template>

<script>
import Balances from '@/components/Balances'
import Delegations from '@/components/Delegations'
import Undelegations from '@/components/Undelegations'
export default {
  name: 'MyNFT',
  components: { Undelegations, Delegations, Balances },
}
</script>

<style scoped>
.my-nft {
  display: flex;
}
</style>
