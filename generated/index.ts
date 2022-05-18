// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import ChronicNetworkChtChronicNetworkChtCht from './ChronicNetwork/cht/ChronicNetwork.cht.cht'
import ChronicNetworkChtdChronicNetworkChtCht from './ChronicNetwork/chtd/ChronicNetwork.cht.cht'


export default { 
  ChronicNetworkChtChronicNetworkChtCht: load(ChronicNetworkChtChronicNetworkChtCht, 'ChronicNetwork.cht.cht'),
  ChronicNetworkChtdChronicNetworkChtCht: load(ChronicNetworkChtdChronicNetworkChtCht, 'ChronicNetwork.cht.cht'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
