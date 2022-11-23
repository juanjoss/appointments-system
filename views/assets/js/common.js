export const toggleSpinner = (name) => {
    const spinner = document.getElementById(name)
    spinner.hidden = !spinner.hidden
}

export const toggleToast = (msg, bgColor) => {
    const toastContainer = document.getElementById("toast")
    toastContainer.className = `toast align-items-center position-fixed top-0 end-0 m-3 border-0 text-bg-${bgColor}`

    const toastBody = document.getElementById("toastBody")
    toastBody.innerHTML = msg

    if (toastContainer) {
        const toast = new bootstrap.Toast(toastContainer)
        toast.show()
    }
}

export async function getCalendarSettings() {
    toggleSpinner("calendarSpinner")

    const response = await fetch(
        "/api/calendar",
        {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json"
            }
        }
    )

    if (!response.ok) {
        document.getElementById("calendarSpinner").className = "spinner-border text-danger"

        const error = document.getElementById("calendarSettingsMsg")
        error.hidden = false

        throw new Error(`HTTP error status: ${response.status}`)
    }

    return await response.json()
}

export async function getAppointmentsFor(date) {
    toggleSpinner("hoursListSpinner")
    document.getElementById("attentionHoursList").innerHTML = ""

    const response = await fetch(
        "/api/attentionHours",
        {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json"
            },
            body: JSON.stringify({
                "year": date.getUTCFullYear(),
                "month": date.getUTCMonth()+1,
                "day": date.getUTCDate()
            })
        }
    )

    if (!response.ok) {
        toggleSpinner("hoursListSpinner")
        toggleToast("Ha ocurrido un error.", "danger")

        throw new Error(`HTTP error status: ${response.status}`)
    }

    return await response.json()
}