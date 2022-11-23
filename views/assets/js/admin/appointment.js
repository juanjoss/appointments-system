import { getCalendarSettings, getAppointmentsFor, toggleSpinner, toggleToast } from "../common.js"

async function updateAppointment(date, weekday, from, to, status) {
    const response = await fetch(
        "/api/appointments",
        {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json"
            },
            body: JSON.stringify({
                "date": date,
                "weekday": weekday,
                "start_hour": from,
                "end_hour": to,
                "status": status
            })
        }
    )

    if (!response.ok) {
        toggleToast("Ha ocurrido un error.", "danger")

        const btn = document.getElementById(`from${from}To${to}`)
        btn.checked = !btn.checked

        throw new Error(`HTTP error status: ${response.status}`)
    }

    return response
}

(() => {
    "use strict"

    getCalendarSettings().then((calendar) => {
        if (calendar.excluded_week_days.length === 7) {
            document.getElementById("notDaysError").hidden = false
        }

        var picker = $('#datepicker').datepicker({
            format: calendar.date_format,
            language: calendar.language,
            daysOfWeekDisabled: calendar.excluded_week_days,
            startDate: new Date(),
            endDate: calendar.end_date,
            maxViewMode: calendar.max_view_mode,
            orientation: "bottom",
            todayHighlight: true
        })

        picker.on("changeDate", (changeDateEvent) => {
            $('#datepicker-input').val(picker.datepicker('getFormattedDate'))

            getAppointmentsFor(changeDateEvent.date).then((data) => {
                const list = document.getElementById("attentionHoursList")

                if (data !== null) {
                    data.forEach(hour => {
                        const div = document.createElement("div")
                        div.className = "form-check form-switch m-3"

                        const input = document.createElement("input")
                        input.className = "form-check-input"
                        input.type = "checkbox"
                        input.id = `from${hour.start_hour}To${hour.end_hour}`
                        input.onclick = (e) => {
                            updateAppointment(
                                changeDateEvent.date,
                                changeDateEvent.date.getUTCDay(),
                                hour.start_hour,
                                hour.end_hour,
                                e.target.checked
                            ).then(() => {
                                toggleToast("Turno actualizado exitosamente.", "success")
                            })
                        }

                        const label = document.createElement("label")
                        label.className = "form-check-label"
                        label.setAttribute("role", "switch")
                        label.htmlFor = `from${hour.start_hour}To${hour.end_hour}`
                        label.textContent = `${hour.start_hour}:00 a ${hour.end_hour}:00`

                        if (hour.status === "available") {
                            input.checked = true
                        } else if (hour.status === "reserved") {
                            input.checked = false
                            input.setAttribute("disabled", true)

                            label.className = "form-check-label text-danger"
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