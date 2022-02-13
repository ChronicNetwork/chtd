<template>
  <div class="swap-container">
    <div class="card-box">
      <div class="swap-background">
        <div class="swap-top">
          <div class="swap-top-sticker" />
          <div class="swap-top-footer" />
        </div>
        <div class="swap-inner-back">
          <div class="swap-setting-box">
            <button class="swap-setting">
              <AccountSettings fill-color="#555e75" />
            </button>
          </div>

          <div class="option-container">
            <div class="option-header-container">
              <p class="option-header-title">From</p>
              <div class="option-header-tailing">
                <p class="option-tailing-item">Available</p>
                <p class="option-tailing-value">0 CHT</p>
                <button class="option-half-btn">MAX</button>
                <button class="option-half-btn">HALF</button>
              </div>
            </div>
            <div class="flex-row">
              <div></div>
            </div>
          </div>
          <div class="option-container">
            <div class="option-header-container">
              <p class="option-header-title">To</p>
            </div>
          </div>
          <div class="invoice-pad">
            <div class="option-header-container">
              <p class="rate-text">Rate</p>
              <p class="rate-text">1 AKT = 0.052 ATOM</p>
            </div>
            <div class="flex-row flex-end">
              <p>1 ATOM = 18.89 AKT</p>
            </div>
            <div class="separator" />
            <div class="">
              <p class="">Estimated Slippage</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AccountSettings from 'vue-material-design-icons/Cog'

export default {
  name: 'Swap',
  components: { AccountSettings },
}
</script>

<style scoped>
.swap-container {
  justify-content: center;
  align-items: center;
  display: flex;
}

.card-box {
  padding: 3rem 2rem;
  width: 100%;
  height: 100%;
  max-width: 520px;
}

.swap-background {
  max-height: 678px;
  width: 100%;
  padding: 10px;
  border: 2px solid rgb(60, 53, 109);
  box-shadow: rgb(0 0 0 / 0%) 0 0 0 0, rgb(0 0 0 / 0%) 0 0 0 0, rgb(0 0 0 / 14%) 0 4px 5px 0,
    rgb(0 0 0 / 12%) 0 1px 10px 0, rgb(0 0 0 / 20%) 0 2px 4px 0;
  background-color: rgb(45, 39, 85);
  border-radius: 1rem;
  position: relative;
}

.swap-inner-back {
  padding: 20px 20px 30px;
  width: 100%;
  height: 100%;
  background-color: rgb(60, 53, 109);
  border-radius: 0.375rem;
}

.swap-top {
  background: linear-gradient(rgb(58, 51, 105) 0%, rgb(35, 29, 75) 100%);
  position: absolute;
  left: 50%;
  top: -8px;
  transform: translate(-50%, 0px);
  box-shadow: rgb(11 16 38 / 48%) 0 2px 2px;
  overflow: hidden;
  z-index: 2;
  border-radius: 0.375rem;
  height: 60px;
  width: 160px;
}

.swap-top-sticker {
  height: 30px;
  width: 48px;
  bottom: 7px;
  left: 50%;
  transform: translate(-50%, 0px);
  background: rgba(91, 83, 147, 0.12);
  background-blend-mode: difference;
  position: absolute;
  border-radius: 0.375rem;
  z-index: 10;
  box-shadow: rgb(0 0 0 / 25%) 1px 1px 1px inset;
}

.swap-top-footer {
  height: 20px;
  position: absolute;
  width: 100%;
  left: 50%;
  bottom: 0;
  transform: translate(-50%, 0px);
  background: linear-gradient(rgb(51, 44, 97) 0%, rgb(49, 42, 93) 10.94%, rgb(45, 39, 85) 100%);
  z-index: 0;
  border-bottom-right-radius: 0.375rem;
  border-bottom-left-radius: 0.375rem;
}

.swap-setting-box {
  width: 100%;
  display: flex;
  -webkit-box-pack: end;
  justify-content: flex-end;
  position: relative;
  z-index: 10;
}

.swap-setting {
  width: 40px;
  height: 40px;
  background-color: transparent;
  border-width: 0;
}

.option-container {
  position: relative;
  border-radius: 1rem;
  padding: 16px 20px 8px 16px;
  background-color: white;
  margin-top: 1.5rem;
  display: flex;
  flex-direction: column;
}

.option-header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.8rem;
  width: 100%;
}

.option-header-title {
  font-size: 16px;
  font-weight: 600;
  margin: 0;
}

.option-header-tailing {
  display: flex;
  flex-direction: row;
  align-items: center;
  height: 100%;
}

.option-tailing-item {
  font-size: 1.175rem;
  font-weight: 600;
  margin: 0 0.7rem 0 0;
}

.option-tailing-value {
  font-size: 1.3rem;
  font-weight: 600;
  margin: 0;
  color: rgb(138, 134, 255);
}

.option-half-btn {
  padding: 4px 6px;
  background-color: rgba(255, 255, 255, 0.12);
  border-radius: 0.375rem;
  margin: 0 0 0 0.7rem;
}

.invoice-pad {
  width: 100%;
  margin: 1.5rem 0;
  border: 1px solid rgba(255, 255, 255, 0.12);
  border-radius: 0.5rem;
  padding: 12px 18px;
  background-color: rgba(255, 255, 255, 0.12);
}

.rate-text {
  font-size: 1.25rem;
  font-weight: 600;
  margin: 0;
}

.separator {
  border-top-width: 1px;
  width: 100%;
  height: 1px;
  background-color: #1a202c;
  opacity: 0.2;
}
</style>
