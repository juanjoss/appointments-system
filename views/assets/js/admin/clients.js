(() => {
    "use strict"

    window.addEventListener("load", () => {
        // birthdate max value
        const birthdate = document.getElementById("birthdate_")
        if (birthdate !== null) {
            const today = new Date()
            const year = today.getUTCFullYear()
            const month = today.getUTCMonth() + 1
            const day = today.getUTCDate()
            
            var maxDate = `${year}-`

            if (month < 10) {
                maxDate += `0${month}-`
            } else {
                maxDate += `${month}-`
            }

            if (day < 10) {
                maxDate += `0${day}`
            } else {
                maxDate += `${day}`
            }

            birthdate.setAttribute("max", maxDate)
            birthdate.addEventListener("change", () => {
                const input = document.getElementById("birthdate")
                input.setAttribute("value", birthdate.value + "T00:00:00Z")
            })

            const input = document.getElementById("birthdate")
            input.setAttribute("value", birthdate.value + "T00:00:00Z")
        }

        const country = $("#country")
        const countryCode = document.getElementById("country_code")
        if (country !== null && countryCode !== null) {
            country.countrySelect({
                defaultCountry: countryCode.value,
                responsiveDropdown: true
            })
        }

        // form validator
        var forms = document.getElementsByClassName('needs-validation')

        Array.prototype.filter.call(forms, (form) => {
            form.addEventListener('submit', (event) => {
                if (form.checkValidity() === false) {
                    event.preventDefault()
                    event.stopPropagation()
                }

                form.classList.add('was-validated')
            }, false)
        })
    })
})()