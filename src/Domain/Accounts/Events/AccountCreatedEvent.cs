using System;
using NutritionTracker.Domain.Common;

namespace NutritionTracker.Domain.Accounts.Events;

public class AccountCreatedEvent : BaseEvent
{
    public Guid AccountId { get; set; }

    public AccountCreatedEvent(Guid accountId)
    {
        AccountId = accountId;
    }
}