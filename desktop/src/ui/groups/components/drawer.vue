<template>
  <v-list nav dense class="text-left">

    <v-list-item exact :to="`/groups/${this.$route.params.groupId}/members`">
      <v-list-item-icon>
        <v-icon>mdi-account-group</v-icon>
      </v-list-item-icon>
      <v-list-item-content>
        <v-list-item-title>Members</v-list-item-title>
      </v-list-item-content>
    </v-list-item>

    <v-list-item exact :to="`/groups/${this.$route.params.groupId}/billing`">
      <v-list-item-icon>
        <v-icon>mdi-credit-card</v-icon>
      </v-list-item-icon>
      <v-list-item-content>
        <v-list-item-title>Billing</v-list-item-title>
      </v-list-item-content>
    </v-list-item>

    <v-list-item exact :to="`/groups/${this.$route.params.groupId}/preferences`">
      <v-list-item-icon>
        <v-icon>mdi-settings</v-icon>
      </v-list-item-icon>
      <v-list-item-content>
        <v-list-item-title>Preferences</v-list-item-title>
      </v-list-item-content>
    </v-list-item>

  </v-list>
</template>


<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class Drawer extends Vue {
  // props
  // data
  // computed
  // lifecycle
  // watch
  // methods
}
</script>


<style lang="scss" scoped>
</style>
