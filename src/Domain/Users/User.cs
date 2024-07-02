using NutritionTracker.Domain.Common;

namespace NutritionTracker.Domain.Users;

public class User : BaseEntity
{
    public string IdentityId { get; set; }

    public User()
    {
        IdentityId = "";
    }
}