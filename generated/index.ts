// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import ChronicNetworkChtChronicNetworkChtCht from './ChronicNetwork/chtd/ChronicNetwork.chtd.cht'
import ChronicNetworkChtdChronicNetworkChtCht from './ChronicNetwork/chtd/ChronicNetwork.chtd.cht'


export default { 
  ChronicNetworkChtChronicNetworkChtCht: load(ChronicNetworkChtChronicNetworkChtCht, 'ChronicNetwork.chtd.cht'),
  ChronicNetworkChtdChronicNetworkChtCht: load(ChronicNetworkChtdChronicNetworkChtCht, 'ChronicNetwork.chtd.cht'),
  
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
