htmx.defineExtension('reset-on-success', {
  onEvent: function (name, event) {
    // console.log('reset-on-success onEvent', {name, event});
    if (name !== 'htmx:beforeOnLoad') return
    if (event.detail.isError) return

    const triggeringElt = event.detail.requestConfig.elt
    if (
      !triggeringElt.closest('[hx-reset-on-success]') &&
      !triggeringElt.closest('[data-hx-reset-on-success]')
    )
      return

    switch (triggeringElt.tagName) {
      case 'INPUT':
      case 'TEXTAREA':
        triggeringElt.value = triggeringElt.defaultValue
        break
      case 'SELECT':
        //too much work
        break
      case 'FORM':
        triggeringElt.reset()
        break
    }
  }
})

console.log('htmxextension.js loaded')
