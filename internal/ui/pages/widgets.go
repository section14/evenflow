package pages

type Widget struct {
    ID string
    Title string
    Value string
}

func GetWidgets(id string) []Widget {
    widgets := []Widget{}
    widgetsGroupA := make([]Widget, 0)
    widgetsGroupB := make([]Widget, 0)

    //test data
    widgetA := Widget{
        ID: "widget_a",
        Title: "Widget A",
        Value: "The value for widget A",
    }

    widgetB := Widget{
        ID: "widget_b",
        Title: "Widget B",
        Value: "The value for widget B. With more text",
    }

    widgetC := Widget{
        ID: "widget_c",
        Title: "Widget C",
        Value: "Widget C value, in a different format",
    }

    widgetD := Widget{
        ID: "widget_d",
        Title: "Widget D",
        Value: "Widget D money",
    }

    widgetE := Widget{
        ID: "widget_e",
        Title: "Widget E",
        Value: "Widget E is for elephant (eats peanuts)",
    }

    widgetF := Widget{
        ID: "widget_f",
        Title: "Widget F",
        Value: "This happens to be the Widget F value",
    }

    widgetsGroupA = append(widgetsGroupA, widgetA, widgetB, widgetC)
    widgetsGroupB = append(widgetsGroupB, widgetD, widgetE, widgetF)
    widgetsGroupB = append(widgetsGroupB, widgetsGroupA...)

    if id == "a" {
        widgets = widgetsGroupA
    } else {
        widgets = widgetsGroupB
    }

    return widgets
}
