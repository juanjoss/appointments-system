import { getCalendarSettings, getAppointmentsFor, toggleSpinner } from "../common.js"

(() => {
    'use strict'

    /*
        Bootstrap Form Validator
    */
    window.addEventListener('load', () => {
        // birthdate max value
        const birthdate = document.getElementById("birthdate_")
        const today = new Date()

        if (birthdate) {
            birthdate.setAttribute("max", `${today.getUTCFullYear()}-0${today.getUTCMonth()}-${today.getUTCDate()}`)
            birthdate.addEventListener("change", () => {
                const input = document.getElementById("birthdate")
                input.value = birthdate.value + "T00:00:00Z"
            })
        }

        // form validator
        var forms = document.getElementsByClassName('needs-validation')

        Array.prototype.filter.call(forms, (form) => {
            const attentionDay = document.getElementById("attentionDay")
            const calendarMsg = document.getElementById("calendarSettingsMsg")

            const attentionHourStart = document.getElementById("attentionHourStart")
            const attentionHourEnd = document.getElementById("attentionHourEnd")
            const appointmentsMsg = document.getElementById("appointmentsMsg")

            form.addEventListener('submit', (event) => {
                if (form.checkValidity() === false) {
                    event.preventDefault()
                    event.stopPropagation()
                }

                if (attentionDay.value === "") {
                    event.preventDefault()
                    event.stopPropagation()

                    calendarMsg.textContent = "Seleccione una fecha para su turno."
                    calendarMsg.hidden = false
                } else {
                    calendarMsg.hidden = true
                }

                if (attentionDay.value !== "" && attentionHourStart.value === "" && attentionHourEnd.value === "") {
                    event.preventDefault()
                    event.stopPropagation()

                    appointmentsMsg.textContent = "Seleccione un horario para su turno."
                    appointmentsMsg.hidden = false
                } else {
                    appointmentsMsg.hidden = true
                }

                form.classList.add('was-validated')
            }, false)
        })
    }, false)

    /*
        Calendar
    */
    const days = ["Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Sábado"]
    const months = ["Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"]

    getCalendarSettings().then((calendar) => {
        if (calendar.excluded_week_days.length === 7) {
            document.getElementById("appointmentInfoTitle").textContent = ""
            document.getElementById("appointmentInfoContent").textContent = "appointments-system/sentimos, no se encuentran turnos disponibles actualmente."
            document.getElementById("appointmentInfo").hidden = false
        }

        var picker = $("#datepicker").datepicker({
            format: calendar.date_format,
            language: calendar.language,
            daysOfWeekDisabled: calendar.excluded_week_days,
            startDate: new Date(),
            endDate: calendar.end_date,
            maxViewMode: calendar.max_view_mode,
            orientation: "bottom",
            todayHighlight: true,
            autoclose: true
        })

        $("#country").countrySelect({
            defaultCountry: "ar",
            responsiveDropdown: true
        })

        picker.on("changeDate", (changeDateEvent) => {
            $("#datepicker-input").val(changeDateEvent.date.toISOString())

            getAppointmentsFor(changeDateEvent.date).then((data) => {
                document.getElementById("attentionDay").setAttribute("value", changeDateEvent.date.getUTCDay())
                document.getElementById("appointmentInfo").hidden = true
                const list = document.getElementById("attentionHoursList")

                if (data !== null) {
                    data.forEach((hour) => {
                        const div = document.createElement("div")
                        div.className = "form-check m-2"

                        const input = document.createElement("input")
                        input.className = "form-check-input"
                        input.type = "radio"
                        input.name = "appointment_hour"
                        input.required = true
                        input.onclick = () => {
                            document.getElementById("attentionHourStart").setAttribute("value", hour.start_hour)
                            document.getElementById("attentionHourEnd").setAttribute("value", hour.end_hour)

                            document.getElementById("appointmentInfoTitle").textContent = "Ha elegido su turno para el día:"

                            const info = document.getElementById("appointmentInfoContent")
                            info.textContent = `
                                ${days[changeDateEvent.date.getUTCDay()]} 
                                ${changeDateEvent.date.getUTCDate()} de ${months[changeDateEvent.date.getUTCMonth()]} del ${changeDateEvent.date.getUTCFullYear()}
                                de ${hour.start_hour} a ${hour.end_hour} horas.
                            `

                            document.getElementById("appointmentInfo").hidden = false
                        }

                        const label = document.createElement("label")
                        label.className = "form-check-label"
                        label.setAttribute("role", "switch")
                        label.htmlFor = `from${hour.start_hour}To${hour.end_hour}`
                        label.textContent = `${hour.start_hour}:00 a ${hour.end_hour}:00`

                        if (hour.status !== "available") {
                            input.setAttribute("disabled", true)

                            if (hour.status === "reserved") {
                                label.classList.add("text-danger")
                            }
                        }

                        div.append(input, label)
                        list.append(div)
                    })
                } else {
                    list.innerHTML = "Turnos no disponibles."
                }

                toggleSpinner("hoursListSpinner")
            })
        })

        toggleSpinner("calendarSpinner")
    })
})()